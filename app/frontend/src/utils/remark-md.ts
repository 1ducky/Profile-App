import { visit } from "unist-util-visit";
import type { Plugin } from "unified";
import type { Root } from "mdast";

interface Options {
    baseURL: string;
}

export const remarkAssetResolver: Plugin<
    [Options],
    Root
> = (options) => {

    // Transformer
    return (tree) => {

        // Traversal AST
        visit(tree, "image", (node) => {

            // Hanya proses URI internal
            if (!node.url.startsWith("asset://")) {
                return;
            }

            // Ambil fileKey
            const fileKey = node.url.slice("asset://".length);

            // Transform URL
            node.url = `${options.baseURL}/assets/${fileKey}`;
        });

    };
};