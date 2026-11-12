
import type { ApiClient } from "../../../lib/api/api-client";
import type { LoginInput, UserSession } from "./types";





export function createAuthApi(api: ApiClient) {
    return {
        login(input: LoginInput) {
            return api.post<UserSession>("/auth/login", input);
        },

        session() {
            return api.get<UserSession>("/auth/session");
        },

        logout() {
            return api.post<null>("/auth/logout");
        },
    };
}