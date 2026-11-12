import { useState } from "react";
import type { ArticleResponse } from "../../../contract/article";
import { FileText, Plus } from "lucide-react";
import { Button } from "../../../components/global/primitive/button";
import { ArticleListItem } from "../../../components/dashboard/Article/ListArticle";
import { DeleteArticles } from "../../../app/feature/article";
import { ApiError } from "../../../lib/api/error";
import { useAlert } from "../../../hooks/useAllert";

export function ArticleManage({ articles }: { articles?: ArticleResponse[] }) {
        const confirm = useAlert();
    const [list, setList] = useState<ArticleResponse[]>(articles || [])
    const onDeleteArticle = async (articleId : string) => {
        try{
            const res = await DeleteArticles(articleId)
            const ok = await confirm({title:"Berhasil DiHapus"})
            setList(list.filter(e => e.id != articleId))
        }catch(e){
            if (e instanceof ApiError) {
                const ok = await confirm({title:`Gagal DI Hapus, Code : ${e.code} ,RequestID: ${e.requestId}`})
            }
            const ok = await confirm({title:`Terjadi Kesalahan Yang Tidak Terduga`})
        }

    }
    return (
        <>
            {list.length == 0 && (
                <EmptyState redirectTo="/dashboard/articles/editor" />
            )}
            {
                list.length >= 1 && (
                    <div className="flex flex-col gap-3 p-3">
                        {
                            list.map((art, i) => {
                                return (
                                    <ArticleListItem article={art} onDelete={(id) => onDeleteArticle(id)} root="/dashboard/articles/editor" />
                                )
                            })
                        }
                    </div>
                )
            }
        </>
    )
}

function EmptyState({ redirectTo }: { redirectTo: string }) {
    return (
        <div className="flex flex-col items-center gap-4 rounded-[16px] border border-dashed border-gray-300 bg-white px-6 py-16 text-center">
            <span className="flex h-14 w-14 items-center justify-center rounded-full bg-[#F5EFE3] text-[#8A6B2F]">
                <FileText className="h-7 w-7" aria-hidden="true" />
            </span>
            <div>
                <h3 className="text-base font-bold">Belum ada konten</h3>
                <p className="mt-1 max-w-sm text-sm text-[#6B7280]">
                    Tulis artikel, tambahkan karya, atau dokumentasikan kegiatan sosial pertama Anda.
                </p>
            </div>
            <a href={redirectTo}>
                <Button icon={<Plus className="h-4 w-4" />}>

                    Buat konten
                </Button>
            </a>
        </div>
    );
}