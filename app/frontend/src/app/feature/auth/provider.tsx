import {
    useCallback,
    useEffect,
    useState,
    type ReactNode,
} from "react";
import type { AuthState } from "./types";
import { ApiError } from "../../../lib/api/error";
import { authApi } from ".";
import { AuthContext } from "./context";

interface Props {
    children: ReactNode;
    initialState?: AuthState;
}

export function AuthProvider({
    children,
    initialState,
}: Props) {
    const [state, setState] = useState<AuthState>(
        initialState ?? {
            status: "loading",
            user: null,
        },
    );

    const refreshSession = useCallback(async () => {
        try {
            const response = await authApi.session();

            setState({
                status: "authenticated",
                user: response.data,
            });
        } catch (error) {
            if (
                error instanceof ApiError &&
                error.statusCode === 401
            ) {
                setState({
                    status: "unauthenticated",
                    user: null,
                });

                return;
            }

            setState({
                status: "error",
                user: null,
                error: error instanceof Error
                    ? error
                    : new Error("Session verification failed"),
            });
        }
    }, []);

    const login = useCallback(
        async (input: Parameters<typeof authApi.login>[0]) => {
            await authApi.login(input);

            await refreshSession();

            // Do not silently treat an unsuccessful
            // session verification as a successful login.
        },
        [refreshSession],
    );

    const logout = useCallback(async () => {
        await authApi.logout();

        setState({
            status: "unauthenticated",
            user: null,
        });
    }, []);

    useEffect(() => {
        if (state.status === "loading") {
            void refreshSession();
        }
    }, [state.status, refreshSession]);

    return (
        <AuthContext.Provider
            value={{
                state,
                login,
                logout,
                refreshSession,
            }}
        >
            {children}
        </AuthContext.Provider>
    );
}