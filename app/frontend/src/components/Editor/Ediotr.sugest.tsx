"use client"; // hapus jika tidak memakai Next.js App Router

import { useId, useRef, useState } from "react";

type Props = {
    name?: string;
    id?: string;
    defaultValue?: string;
    suggestions: string[];
    placeholder?: string;
    fieldRef: React.RefObject<HTMLInputElement | null>
};

export function SuggestionsField({
    name = "category",
    id = "category",
    defaultValue = "",
    suggestions,
    placeholder,
    fieldRef
}: Props) {

    const listId = useId();

    const [query, setQuery] = useState(defaultValue); // hanya untuk filter saran
    const [open, setOpen] = useState(false);
    const [active, setActive] = useState(-1);

    const filtered = suggestions.filter(
        (s) =>
            s.toLowerCase().includes(query.trim().toLowerCase()) &&
            s.toLowerCase() !== query.trim().toLowerCase()
    );

    const select = (value: string) => {
        if (fieldRef.current) fieldRef.current.value = value; // ubah value uncontrolled
        setQuery(value);
        setOpen(false);
        setActive(-1);
        fieldRef.current?.focus();
    };

    const onKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (!open || filtered.length === 0) return;

        if (e.key === "ArrowDown") {
            e.preventDefault();
            setActive((i) => (i + 1) % filtered.length);
        } else if (e.key === "ArrowUp") {
            e.preventDefault();
            setActive((i) => (i <= 0 ? filtered.length - 1 : i - 1));
        } else if (e.key === "Enter" && active >= 0) {
            e.preventDefault(); // cegah submit form
            select(filtered[active]);
        } else if (e.key === "Escape") {
            setOpen(false);
        }
    };

    return (
        <div className="relative">
            <input
                ref={fieldRef}
                id={id}
                name={name}
                defaultValue={defaultValue}
                placeholder={placeholder}
                autoComplete="off"
                role="combobox"
                aria-expanded={open && filtered.length > 0}
                aria-controls={listId}
                aria-autocomplete="list"
                onChange={(e) => {
                    setQuery(e.target.value);
                    setOpen(true);
                    setActive(-1);
                }}
                onFocus={() => setOpen(true)}
                onBlur={() => setOpen(false)}
                onKeyDown={onKeyDown}
                className="w-full rounded-[10px] border border-[#D1D5DB] bg-white px-3.5 py-2.5 text-[15px] text-[#1F2937] placeholder:text-[#9CA3AF] outline-none transition focus:border-primary focus:ring-2 focus:ring-primary/20"
            />

            {open && filtered.length > 0 && (
                <ul
                    id={listId}
                    role="listbox"
                    className="absolute z-20 mt-1.5 max-h-56 w-full overflow-auto rounded-[12px] border border-[#E5E7EB] bg-white py-1 shadow-[0_4px_16px_rgba(31,41,55,0.06)]"
                >
                    {filtered.map((s, i) => (
                        <li
                            key={s}
                            role="option"
                            aria-selected={i === active}
                            // onMouseDown + preventDefault agar input tidak blur sebelum klik terproses
                            onMouseDown={(e) => {
                                e.preventDefault();
                                select(s);
                            }}
                            onMouseEnter={() => setActive(i)}
                            className={`cursor-pointer px-3.5 py-2 text-[14px] text-[#1F2937] ${i === active ? "bg-primary/10 text-primary" : ""
                                }`}
                        >
                            {s}
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
}