import { useImperativeHandle, useRef, useState, type Ref } from "react";

export type TagInputHandle = {
    getValue: () => string[];
    clear: () => void;
};

type TagInputProps = {
    id: string;
    placeholder?: string;
    defaultValue?: string[];
    ref?: Ref<TagInputHandle>; // React 19: ref sudah jadi prop biasa
};

export function TagInput({ id, placeholder, defaultValue = [], ref }: TagInputProps) {
    const [draft, setDraft] = useState("");
    const [tags, setTags] = useState<string[]>(defaultValue);

    // Mirror terbaru, supaya getValue() selalu akurat tanpa stale closure
    const tagsRef = useRef<string[]>(defaultValue);

    const update = (next: string[]) => {
        tagsRef.current = next;
        setTags(next);
    };

    useImperativeHandle(ref, () => ({
        getValue: () => tagsRef.current,
        clear: () => update([]),
    }));

    const commit = () => {
        const t = draft.trim();
        if (t && !tagsRef.current.includes(t)) update([...tagsRef.current, t]);
        setDraft("");
    };

    const onDelete = (item: string) => {
        update(tagsRef.current.filter((t) => t !== item));
    };

    return (
        <div className="flex flex-wrap items-center gap-2 rounded-[10px] border border-[#D1D5DB] bg-white px-2.5 py-2 focus-within:border-primary focus-within:ring-2 focus-within:ring-primary/20">
            {tags.map((t) => (
                <span
                    key={t}
                    className="inline-flex items-center gap-1 rounded-full bg-section-warm px-2.5 py-1 text-sm text-[#1F2937]"
                >
                    {t}
                    <button
                        type="button"
                        aria-label={`Hapus ${t}`}
                        onClick={() => onDelete(t)}
                        className="text-text-muted hover:text-[#1F2937]"
                    >
                        ×
                    </button>
                </span>
            ))}
            <input
                id={id}
                value={draft}
                placeholder={tags.length ? "" : placeholder}
                onChange={(e) => setDraft(e.target.value)}
                onKeyDown={(e) => {
                    if (e.key === "Enter" || e.key === ",") {
                        e.preventDefault();
                        commit();
                    } else if (e.key === "Backspace" && !draft && tagsRef.current.length) {
                        update(tagsRef.current.slice(0, -1));
                    }
                }}
                onBlur={commit}
                className="min-w-30 flex-1 bg-transparent px-1 py-1 text-[15px] text-[#1F2937] outline-none placeholder:text-[#9CA3AF]"
            />
        </div>
    );
}