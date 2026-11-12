import { PublishArticle } from "../../../../app/feature/article"
import { EditorArticle } from "../../../../components/Editor/EditorArticle"
import type { ArticleField } from "../../../../contract/article"


export function CreateArticle() {
    const onSubmit = async (req: ArticleField) => {
        try {
            const res = await PublishArticle(req)
            console.log(res)
            window.location.href = "/dashboard/articles"
        } catch (e) {
            console.log(e)
        }
    }

    return <EditorArticle
        onPublish={(e) => onSubmit(e)}
    />
}