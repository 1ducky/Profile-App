export type Status = "success" | "error";

export interface ApiResponse<T> {
    status: Status;
    data: T;
}

export interface ErrorResponse {
    code: string;
    message: string;
    request_id: string;
    status: Status;
}