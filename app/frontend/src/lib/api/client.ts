import { ApiError } from "./error";
import type { ApiResponse, ErrorResponse } from "./types";


interface RequestOptions extends Omit<RequestInit, "body"> {
    body?: unknown;
    cookie? : string
}

const isServer = typeof window === "undefined";

function getApiBase(): string {
    if (!isServer) {
        return "/api";
    }

    const base = process.env.API_URL;

    if (!base) {
        throw new Error("API_URL is not configured");
    }

    return base.replace(/\/+$/, "");
}

async function request<T>(
    path: string,
    options: RequestOptions = {},
): Promise<ApiResponse<T>> {
    const {
        cookie,
        body,
        headers: customHeaders,
        ...init
    } = options;

    const headers = new Headers(customHeaders);

    headers.set("Accept", "application/json");

    let requestBody: BodyInit | undefined;

    if (cookie) {
        headers.set("Cookie", cookie);
        console.log("Cookie Attcatceh")
        console.log(cookie)
    }

    if (body instanceof FormData) {
        requestBody = body;
    } else if (body !== undefined) {
        headers.set("Content-Type", "application/json");
        requestBody = JSON.stringify(body);
    }

    const response = await fetch(
        `${getApiBase()}/${path.replace(/^\/+/, "")}`,
        {
            ...init,
            headers,
            body: requestBody,
            credentials: "same-origin",
        },
    );

    const raw = await response.text();

    let payload: unknown;

    try {
        payload = raw ? JSON.parse(raw) : null;
    } catch {
        throw new ApiError("INVALID_RESPONSE", "Invalid Response Json", "", 500);
    }

    if (!response.ok) {
        throw ApiError.fromResponse(
            payload as ErrorResponse,
            response.status
        );
    }

    return payload as ApiResponse<T>;
}

export const api = {
    request,

    get<T>(path: string, cookie?:string) {
        
        return request<T>(path, { cookie,method: "GET" });
    },

    post<T>(path: string, body?: unknown, cookie?:string) {
        return request<T>(path, {
            cookie,
            method: "POST",
            body,
        });
    },

    patch<T>(path: string, body?: unknown, cookie?:string) {
        return request<T>(path, {
            cookie,
            method: "PATCH",
            body,
        });
    },

    delete<T>(path: string, cookie?:string) {
        return request<T>(path, {
            cookie,
            method: "DELETE",
        });
    },
};