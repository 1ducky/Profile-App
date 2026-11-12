import { useRef, useState } from "react";
import EditorMarkdown from "./EditorMarkDown";
import type { MDXEditorMethods } from "@mdxeditor/editor";

export default function Editor() {
    const [md, setMD] = useState<string>("")

    const editorRef = useRef<MDXEditorMethods>(null)
    return (
        <>
            <EditorMarkdown markdown={md} ref={editorRef} />
            <button type="button" className="rounded-card border border-primary bg-primary px-5 py-2.5 text-sm font-medium transition-colors focus-visible:outline-3 focus-visible:outline-offset-2 focus-visible:outline-primary/50 text-white" onClick={() => console.log(md)}>Submit</button>
        </>
    )
}