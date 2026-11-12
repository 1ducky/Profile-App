// entity 
export interface UserSession {
    id: string;
    email: string;
    role: string;
    version: number;
}
//input

export interface LoginInput {
    email: string;
    password: string;
}

export type AuthState =
    | {
        status: "loading";
        user: null;
    }
    | {
        status: "authenticated";
        user: UserSession;
    }
    | {
        status: "unauthenticated";
        user: null;
    }
    | {
        status: "error";
        user: null;
        error: Error;
    };