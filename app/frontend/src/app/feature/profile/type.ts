export interface ProfileResponse {
    username: string,
    tag: string,
    avatar: string,
    bio: string,
}
export interface UpdateProfileRequest {
    username: string,
    tag: string,
    bio: string,
}