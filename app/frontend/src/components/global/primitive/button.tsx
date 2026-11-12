type ButtonProps = {
    variant?: "primary" | "secondary" | "ghost";
    icon?: React.ReactNode;
    children?: React.ReactNode;
    className?: string;
} & React.ButtonHTMLAttributes<HTMLButtonElement>;

export function Button({
    variant = "primary",
    icon,
    children,
    className = "",
    ...props
}: ButtonProps) {
    const styles = {
        primary: "bg-[#0B63CE] text-white hover:bg-[#0A55B0] border border-transparent",
        secondary:
            "bg-white text-[#1F2937] border border-gray-300 hover:bg-gray-50",
        ghost: "bg-transparent text-[#1F2937] hover:bg-gray-100 border border-transparent",
    }[variant];

    return (
        <button
            {...props}
            className={`inline-flex items-center justify-center gap-2 rounded-full px-5 py-2.5 text-sm font-semibold transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE] focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 ${styles} ${className}`}
        >
            {icon}
            {children}
        </button>
    );
}

export function IconButton({
    label,
    children,
    className = "",
    ...props
}: { label: string; children: React.ReactNode; className?: string } & React.ButtonHTMLAttributes<HTMLButtonElement>) {
    return (
        <button
            {...props}
            aria-label={label}
            title={label}
            className={`inline-flex h-10 w-10 items-center justify-center rounded-full text-[#6B7280] transition-colors hover:bg-gray-100 hover:text-[#1F2937] focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE] ${className}`}
        >
            {children}
        </button>
    );
}