
import type { LoginInput } from "../../app/feature/auth/types";
import { GetListArticle } from "../../app/feature/article";

import { useRef, useState } from "react";
import { Mail, Lock, Eye, EyeOff, ArrowRight } from "lucide-react";
import { createAuthApi } from "../../app/feature/auth";
import { api } from "../../lib/api/api-client";



export default function LoginForm({
    onLogin,
}: {
    onLogin: (req: LoginInput) => void;
}) {
    const emailRef = useRef<HTMLInputElement>(null);
    const passwordRef = useRef<HTMLInputElement>(null);
    const [showPassword, setShowPassword] = useState(false);

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const email = emailRef.current?.value.trim();
        const password = passwordRef.current?.value;

        if (!email || !password) return;

        const payload = { email, password };

        onLogin(payload);

        // Contoh integrasi API:
        // await authService.login(payload);
    };

    return (
        <form
            onSubmit={handleSubmit}
            method="post"
            className="mx-auto w-full max-w-md space-y-6 rounded-[16px] border border-gray-200 bg-white p-8 shadow-[0_4px_16px_rgba(31,41,55,0.06)] sm:p-10"
        >
            <div className="space-y-2">
                <span className="text-xs font-semibold uppercase tracking-widest text-[#0B63CE]">
                    Selamat Datang
                </span>
                <h1 className="text-2xl font-bold tracking-tight text-[#1F2937]">
                    Masuk ke Akun Anda
                </h1>
                <p className="text-sm leading-relaxed text-[#6B7280]">
                    Silakan masuk untuk melanjutkan.
                </p>
            </div>

            <div className="space-y-2">
                <label
                    htmlFor="email"
                    className="block text-sm font-medium text-[#1F2937]"
                >
                    Email
                </label>
                <div className="relative">
                    <Mail
                        aria-hidden="true"
                        className="pointer-events-none absolute left-4 top-1/2 size-4 -translate-y-1/2 text-[#6B7280]"
                    />
                    <input
                        ref={emailRef}
                        type="email"
                        id="email"
                        name="email"
                        placeholder="you@example.com"
                        autoComplete="email"
                        required
                        className="w-full rounded-full border border-gray-300 bg-[#F8FAFC] py-3 pl-11 pr-4 text-sm text-[#1F2937] placeholder:text-[#6B7280]/70 transition focus:border-[#0B63CE] focus:bg-white focus:outline-none focus:ring-2 focus:ring-[#0B63CE]/30"
                    />
                </div>
            </div>

            <div className="space-y-2">
                <label
                    htmlFor="password"
                    className="block text-sm font-medium text-[#1F2937]"
                >
                    Password
                </label>
                <div className="relative">
                    <Lock
                        aria-hidden="true"
                        className="pointer-events-none absolute left-4 top-1/2 size-4 -translate-y-1/2 text-[#6B7280]"
                    />
                    <input
                        ref={passwordRef}
                        type={showPassword ? "text" : "password"}
                        id="password"
                        name="password"
                        placeholder="Enter your password"
                        autoComplete="current-password"
                        required
                        className="w-full rounded-full border border-gray-300 bg-[#F8FAFC] py-3 pl-11 pr-12 text-sm text-[#1F2937] placeholder:text-[#6B7280]/70 transition focus:border-[#0B63CE] focus:bg-white focus:outline-none focus:ring-2 focus:ring-[#0B63CE]/30"
                    />
                    <button
                        type="button"
                        onClick={() => setShowPassword((prev) => !prev)}
                        aria-label={showPassword ? "Sembunyikan password" : "Tampilkan password"}
                        aria-pressed={showPassword}
                        className="absolute right-3 top-1/2 -translate-y-1/2 rounded-full p-1.5 text-[#6B7280] transition hover:text-[#1F2937] focus:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE]/40"
                    >
                        {showPassword ? (
                            <EyeOff aria-hidden="true" className="size-4" />
                        ) : (
                            <Eye aria-hidden="true" className="size-4" />
                        )}
                    </button>
                </div>
            </div>

            <button
                type="submit"
                className="group inline-flex w-full items-center justify-center gap-2 rounded-full bg-[#0B63CE] px-6 py-3 text-sm font-semibold text-white transition hover:bg-[#0952a8] focus:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE]/50 focus-visible:ring-offset-2 motion-reduce:transition-none"
            >
                Login
                <ArrowRight
                    aria-hidden="true"
                    className="size-4 transition-transform group-hover:translate-x-0.5 motion-reduce:transition-none"
                />
            </button>
        </form>
    );
}

export function Login() {
    const onLogin = async (req: LoginInput) => {
        try {

            const res = await createAuthApi(api).login(req)
            console.log(res)
            window.location.href = "/dashboard/articles"
        } catch (e) {
            console.log(e)
        }

    }

    return <LoginForm onLogin={onLogin} />
}