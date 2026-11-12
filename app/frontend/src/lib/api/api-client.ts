
import { ApiError } from "./error";
import type { ApiResponse, ErrorResponse } from "./types";

interface RequestOptions extends Omit<RequestInit, "body"> {
    body?: unknown;
}

interface ApiContext {
    cookie?: string;
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

function createApi(context: ApiContext = {}) {
    const base = getApiBase();

    async function request<T>(
        path: string,
        options: RequestOptions = {},
    ): Promise<ApiResponse<T>> {
        const {
            body,
            headers: customHeaders,
            ...init
        } = options;

        const headers = new Headers(customHeaders);

        headers.set("Accept", "application/json");

        // Forward cookie only from server-side context.
        if (isServer && context.cookie) {
            headers.set("Cookie", context.cookie);
        }

        let requestBody: BodyInit | undefined;

        if (body instanceof FormData) {
            requestBody = body;
        } else if (body !== undefined) {
            headers.set("Content-Type", "application/json");
            requestBody = JSON.stringify(body);
        }

        const url = `${base}/${path.replace(/^\/+/, "")}`;

        const response = await fetch(url, {
            ...init,
            headers,
            body: requestBody,
            credentials: "same-origin",
        });

        const raw = await response.text();

        let payload: unknown;

        try {
            payload = raw ? JSON.parse(raw) : null;
        } catch {
            throw new ApiError(
                "INVALID_RESPONSE",
                "Invalid Response Json",
                "",
                500,
            );
        }

        if (!response.ok) {
            throw ApiError.fromResponse(
                payload as ErrorResponse,
                response.status,
            );
        }

        return payload as ApiResponse<T>;
    }

    return {
        request,

        get<T>(path: string, options?: RequestOptions) {
            return request<T>(path, {
                ...options,
                method: "GET",
            });
        },

        post<T>(
            path: string,
            body?: unknown,
            options?: RequestOptions,
        ) {
            return request<T>(path, {
                ...options,
                method: "POST",
                body,
            });
        },

        patch<T>(
            path: string,
            body?: unknown,
            options?: RequestOptions,
        ) {
            return request<T>(path, {
                ...options,
                method: "PATCH",
                body,
            });
        },

        delete<T>(path: string, options?: RequestOptions) {
            return request<T>(path, {
                ...options,
                method: "DELETE",
            });
        },
    };
}

// Default instance for browser-side usage.
export const api = createApi();

// Factory for request-scoped SSR usage.
export { createApi };

export interface ApiClient {
    request<T>(
        path: string,
        options?: RequestOptions,
    ): Promise<ApiResponse<T>>;

    get<T>(path: string, options?: RequestOptions):
        Promise<ApiResponse<T>>;

    post<T>(
        path: string,
        body?: unknown,
        options?: RequestOptions,
    ): Promise<ApiResponse<T>>;

    patch<T>(
        path: string,
        body?: unknown,
        options?: RequestOptions,
    ): Promise<ApiResponse<T>>;

    delete<T>(
        path: string,
        options?: RequestOptions,
    ): Promise<ApiResponse<T>>;
}