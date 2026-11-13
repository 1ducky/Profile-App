import { useState } from "react";
import type { ArticleResponse } from "../../../contract/article";
import { FileText, Loader2, Plus, Tag, Trash2 } from "lucide-react";
import { Button } from "../../../components/global/primitive/button";
import { ArticleListItem } from "../../../components/dashboard/Article/ListArticle";
import { DeleteArticles } from "../../../app/feature/article";
import { ApiError } from "../../../lib/api/error";
import { useAlert } from "../../../hooks/useAllert";
import type { categoryResponse } from "../../../app/feature/category/type";
import { createCategoryApi } from "../../../app/feature/category";
import { api } from "../../../lib/api/api-client";

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

type CategoryManagerProps = {
  /** Data awal (opsional), mis. dari server */
  initialCategories?: categoryResponse[];
  /** Endpoint POST untuk menambah kategori */
  /**
   * Dipanggil saat tombol hapus ditekan.
   * Parent yang bertanggung jawab memanggil API delete.
   * Item akan dihapus dari list lokal setelah promise selesai tanpa error.
   */
};
 
export default function CategoryManager({
  initialCategories = [],
}: CategoryManagerProps) {
    const alert = useAlert()
  const [title, setTitle] = useState("");
  const [categories, setCategories] = useState<categoryResponse[]>(initialCategories);
  const [isAdding, setIsAdding] = useState(false);
  const [deletingId, setDeletingId] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
 
  // POST ke API, jika sukses tambahkan ke list lokal
  const onAdd = async () => {
    const value = title.trim();
    if (!value || isAdding) return;
 
    setIsAdding(true);
    setError(null);
 
    try {
      const res = await createCategoryApi(api).createCategory(title) 
      // API diharapkan mengembalikan { id, slug }
 
      setCategories((prev) => [...prev, { id: res.data.id, slug: res.data.slug }]);
      setTitle("");
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : "Terjadi kesalahan. Coba lagi beberapa saat."
      );
    } finally {
      setIsAdding(false);
    }
  };
 
  const handleDelete = async (category: categoryResponse) => {
    setDeletingId(category.id);
    setError(null);
 
    try {
        const alern = await alert({title:"yakin ingin menghapus Category "+category.slug +" ?"})
        const res = await createCategoryApi(api).deleteCategory(category.id)
        const info = await alert({title:"Category "+category.slug +" berhasil di hapus ?"})
        setCategories((prev) => prev.filter((c) => c.id !== category.id));
    } catch(e) {
        if (e instanceof ApiError){
            const alern = await alert({title:`Gagal Menghapus Category CODE:${e.code}`})
            
        }
      setError("Kategori gagal dihapus. Coba lagi.");
    } finally {
      setDeletingId(null);
    }
  };
 
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onAdd();
  };
 
  return (
    <section
      className="w-full  rounded-2xl border border-slate-200 bg-white p-6 sm:p-8"
      style={{ boxShadow: "0 4px 16px rgba(31, 41, 55, 0.06)" }}
      aria-labelledby="category-manager-title"
    >
      <header className="mb-6">
        <h2
          id="category-manager-title"
          className="text-xl font-semibold tracking-tight text-[#1F2937]"
        >
          Kelola kategori
        </h2>
        <p className="mt-1 text-sm leading-relaxed text-[#6B7280]">
          Tambahkan kategori untuk mengelompokkan artikel dan kegiatan.
        </p>
      </header>
 
      {/* Form tambah */}
      <form onSubmit={handleSubmit} className="flex flex-col gap-3 sm:flex-row">
        <label htmlFor="category-title" className="sr-only">
          Nama kategori
        </label>
        <input
          id="category-title"
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="Nama kategori, mis. Literasi Digital"
          disabled={isAdding}
          className="h-11 flex-1 rounded-full border border-slate-300 bg-white px-5 text-sm text-[#1F2937] placeholder:text-[#6B7280] focus:border-[#0B63CE] focus:outline-none focus:ring-2 focus:ring-[#0B63CE]/30 disabled:opacity-60"
        />
        <button
          type="submit"
          disabled={isAdding || !title.trim()}
          className="inline-flex h-11 items-center justify-center gap-2 rounded-full bg-[#0B63CE] px-6 text-sm font-medium text-white transition-colors hover:bg-[#0952a8] focus:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE] focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
        >
          {isAdding ? (
            <Loader2 className="h-4 w-4 animate-spin" aria-hidden />
          ) : (
            <Plus className="h-4 w-4" aria-hidden />
          )}
          Tambah
        </button>
      </form>
 
      {error && (
        <p
          role="alert"
          className="mt-3 rounded-xl bg-red-50 px-4 py-2.5 text-sm text-red-700"
        >
          {error}
        </p>
      )}
 
      {/* List kategori */}
      {categories.length === 0 ? (
        <div className="mt-6 flex flex-col items-center gap-2 rounded-xl bg-[#F5EFE3] px-4 py-8 text-center">
          <Tag className="h-5 w-5 text-[#6B7280]" aria-hidden />
          <p className="text-sm text-[#6B7280]">
            Belum ada kategori. Tulis nama di atas untuk menambahkan yang
            pertama.
          </p>
        </div>
      ) : (
        <ul className="mt-6 flex flex-col gap-2">
          {categories.map((category) => (
            <li
              key={category.id}
              className="flex items-center justify-between gap-3 rounded-xl border border-slate-200 bg-[#F8FAFC] px-4 py-3"
            >
              <div className="flex min-w-0 items-center gap-3">
                <span className="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-[#16A34A]/10 text-[#16A34A]">
                  <Tag className="h-4 w-4" aria-hidden />
                </span>
                <span className="truncate text-sm font-medium text-[#1F2937]">
                  {category.slug}
                </span>
              </div>
 
              <button
                type="button"
                onClick={() => handleDelete(category)}
                disabled={deletingId === category.id}
                aria-label={`Hapus kategori ${category.slug}`}
                className="inline-flex h-9 w-9 shrink-0 items-center justify-center rounded-full text-[#6B7280] transition-colors hover:bg-red-50 hover:text-red-600 focus:outline-none focus-visible:ring-2 focus-visible:ring-red-500 disabled:opacity-50"
              >
                {deletingId === category.id ? (
                  <Loader2 className="h-4 w-4 animate-spin" aria-hidden />
                ) : (
                  <Trash2 className="h-4 w-4" aria-hidden />
                )}
              </button>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
 