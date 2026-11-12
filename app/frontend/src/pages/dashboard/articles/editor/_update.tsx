import { PublishArticle, UpdateArticles } from "../../../../app/feature/article"
import { EditorArticle } from "../../../../components/Editor/EditorArticle"
import type { ArticleField, ArticleResponse } from "../../../../contract/article"


export function UpdateArticleEditor({ defaultValue }: { defaultValue: ArticleResponse }) {
    const onSubmit = async (req: ArticleField) => {
        try {
            const res = await UpdateArticles(req, defaultValue.id)
            console.log(res)
            window.location.href = "/dashboard/articles"
        } catch (e) {
            console.log(e)
        }
    }

    return <EditorArticle
        onPublish={(e) => onSubmit(e)} defaultValue={{
            badges: defaultValue?.badges,
            category: defaultValue?.category,
            content: defaultValue?.content || "",
            cover: defaultValue?.cover,
            excerpt: defaultValue?.excerpt,
            tags: defaultValue?.tags,
            title: defaultValue.title,
            url: defaultValue?.url
        }}
    />
}