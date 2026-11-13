import type { ApiClient } from "../../../lib/api/api-client";
import type { ProfileResponse } from "../profile/type";
import type { categoryResponse } from "./type";

export function createCategoryApi(api: ApiClient) {
    return {
        getCategory() {
            return api.get<categoryResponse[]>("/category");
        },
        createCategory(slug:string) {
            return api.post<categoryResponse>("/category",{slug:slug});
        },
        deleteCategory(id:string) {
            return api.delete<{message:string}>("/category/"+id);
        },

    };
}