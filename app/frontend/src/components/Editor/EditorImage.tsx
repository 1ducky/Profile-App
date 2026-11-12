import { usePublisher, insertImage$, Button } from "@mdxeditor/editor";
import { useImageUploader } from "./ImageUploader/ImageProvider";
import type { ArticleImage } from "../../types/article/entity";
import { Image } from "lucide-react";

export function CustomImageButton() {
    const insertImage = usePublisher(insertImage$);
    const { open } = useImageUploader();

    async function handleClick() {
        const img = await open({ title: "Sisipkan gambar" });
        if (!img) return; // pengguna menutup dialog tanpa memilih

        insertImage({
            src: img.url,      // URL final dari backend (atau URL eksternal)
            altText: img.alt,
            title: "",
        });
    }

    return (
        <Button title="Upload gambar" onClick={handleClick}>
            <Image />
        </Button>
    );
}

export function InsertImageButton({ onSuccess }: { onSuccess: (imageUrl: ArticleImage) => void }) {

    const { open } = useImageUploader();

    async function handleClick() {
        const img = await open({ title: "Sisipkan gambar" });
        if (!img) return; // pengguna menutup dialog tanpa memilih

        onSuccess({
            src: img.url,      // URL final dari backend (atau URL eksternal)
            alt: img.alt,
            caption: img.title,
        });
    }

    return (
        <button type="button" title="Upload gambar" onClick={handleClick} className="flex flex-row flex-nowrap gap-3">
            <Image /> <p>
                Upload Gambar
            </p>
        </button>
    );
}