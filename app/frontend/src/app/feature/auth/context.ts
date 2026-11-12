import { createContext } from "react";
import type { AuthState, LoginInput } from "./types";


export interface AuthContextValue {
    state: AuthState;

    login(input: LoginInput): Promise<void>;
    logout(): Promise<void>;
    refreshSession(): Promise<void>;
}

export const AuthContext =
    createContext<AuthContextValue | null>(null);