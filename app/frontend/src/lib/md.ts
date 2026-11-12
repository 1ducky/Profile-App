
import MarkdownIt from "markdown-it";
import { ResolveUrlAsset } from "../utils/resolve-url";

const markdown = new MarkdownIt({
    html: false,
    linkify: true,
    typographer: true,
    breaks: false,
});

// Tambahkan atribut aman untuk tautan eksternal.
markdown.renderer.rules.link_open = (
    tokens,
    idx,
    options,
    env,
    self,
) => {
    const token = tokens[idx];
    const href = token.attrGet("href") as string;

    if (href && /^https?:\/\//i.test(href)) {
        token.attrSet("target", "_blank");
        token.attrSet("rel", "noopener noreferrer");
    }

    return self.renderToken(tokens, idx, options);
};

// Gambar: resolve asset:... menjadi URL lengkap
markdown.renderer.rules.image = (tokens, idx, options, env, self) => {
    const token = tokens[idx];
    const src = token.attrGet("src") ?? "";

    token.attrSet("src", ResolveUrlAsset(src as string));
    token.attrSet("loading", "lazy");
    token.attrSet("decoding", "async");

    // markdown-it menyimpan alt sebagai children, bukan atribut
    token.attrSet(
        "alt",
        self.renderInlineAsText(token.children ?? [], options, env),
    );

    return self.renderToken(tokens, idx, options);
};


export function renderMarkdown(content: string): string {
    return markdown.render(content);
}
