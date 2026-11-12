import type { APIRoute } from "astro";

const API_URL = process.env.API_URL || "";
const APP_ORIGIN = process.env.APP_ORIGIN || "";

const ALLOWED_METHODS = new Set([
    "GET",
    "POST",
    "PUT",
    "PATCH",
    "DELETE",
]);

export const ALL: APIRoute = async ({ request }) => {
    const method = request.method;
    console.log("Backend",API_URL)
    console.log("Frontend",APP_ORIGIN)
    console.log("fullpath",API_URL+request.url)

    if (!ALLOWED_METHODS.has(method)) {
        return new Response(null, {
            status: 405,
            headers: {
                Allow: [...ALLOWED_METHODS].join(", "),
            },
        });
    }

    // Basic CSRF protection for state-changing requests.
    if (!["GET", "HEAD"].includes(method)) {
        const origin = request.headers.get("origin");

        if (!origin || origin !== APP_ORIGIN) {
            console.log(origin)
            console.log(APP_ORIGIN)
            return new Response("Forbidden origin", {
                status: 403,
            });
        }
    }

    const incomingURL = new URL(request.url);

    const path = incomingURL.pathname.replace(/^\/api\/?/, "");

    const targetURL =
        `${API_URL.replace(/\/+$/, "")}/${path}${incomingURL.search}`;
    const headers = new Headers();

    // Forward only required headers.
    for (const key of [
        "accept",
        "content-type",
        "cookie",
        "x-request-id",
        "if-match",
        "if-none-match",
    ]) {
        const value = request.headers.get(key);

        if (value) {
            headers.set(key, value);
        }
    }

    // Preserve the original browser origin for backend CSRF validation.
    const origin = request.headers.get("origin");

    if (origin) {
        headers.set("origin", origin);
    }

    const body = ["GET", "HEAD"].includes(method)
        ? undefined
        : await request.arrayBuffer();

    let upstream: Response;

    try {
        upstream = await fetch(targetURL, {
            method,
            headers,
            body,
            redirect: "manual",
        });
    } catch {
        return Response.json(
            {
                code: "UPSTREAM_UNAVAILABLE",
                message: "Backend service is unavailable",
                request_id: targetURL,
                status: upstreamStatusPlaceholder(),
            },
            { status: 502 },
        );
    }

    const responseHeaders = new Headers();

    for (const key of [
        "content-type",
        "cache-control",
        "etag",
        "vary",
        "x-request-id",
    ]) {
        const value = upstream.headers.get(key);

        if (value) {
            responseHeaders.set(key, value);
        }
    }

    // Preserve every Set-Cookie header independently.
    const setCookies = upstream.headers.getSetCookie?.() ?? [];

    for (const cookie of setCookies) {
        responseHeaders.append("Set-Cookie", cookie);
    }

    return new Response(upstream.body, {
        status: upstream.status,
        headers: responseHeaders,
    });
};

function upstreamStatusPlaceholder() {
    return "ERROR";
}