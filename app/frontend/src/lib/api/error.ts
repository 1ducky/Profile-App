import type { ErrorResponse } from "./types";



export class ApiError extends Error {
  constructor(
    public readonly code: string,
    message: string,
    public readonly requestId: string,
    public readonly statusCode: number,
  ) {
    super(message);
    this.name = "ApiError";
  }

  static fromResponse(
    response: ErrorResponse,
    statusCode: number,
  ) {
    return new ApiError(
      response.code,
      response.message,
      response.request_id,
      statusCode,
    );
  }
}