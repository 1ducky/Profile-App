import type { ApiClient } from "../../../lib/api/api-client";
import type { ProfileResponse } from "./type";

export function createProfileApi(api: ApiClient) {
    return {
        getProfile() {
            return api.get<ProfileResponse>("/profile");
        },
        getProfileByID(id:string) {
            return api.get<ProfileResponse>("/profile/"+id);
        },
        updateProfile() {
            return api.get<ProfileResponse>("/profile");
        },
        updateAvatar() {
            return api.get<ProfileResponse>("/profile/avatar");
        },

    };
}