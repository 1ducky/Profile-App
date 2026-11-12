import type { ApiClient } from "../../../lib/api/api-client";
import type { ProfileResponse } from "./type";

export function createProfileApi(api: ApiClient) {
    return {
        getProfile() {
            return api.get<ProfileResponse>("/profile");
        },

    };
}