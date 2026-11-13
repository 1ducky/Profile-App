import type { ApiClient } from "../../../lib/api/api-client";
import type { ProfileResponse, UpdateProfileRequest } from "./type";

export function createProfileApi(api: ApiClient) {
    return {
        getProfile() {
            return api.get<ProfileResponse>("/profile");
        },
        getProfileByID(id:string) {
            return api.get<ProfileResponse>("/profile/"+id);
        },
        updateProfile(req : UpdateProfileRequest) {
            return api.patch<ProfileResponse>("/profile",req);
        },
        updateAvatar(filekey:string) {
            return api.post<ProfileResponse>("/profile/avatar",{filekey});
        },

    };
}