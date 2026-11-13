import { useRef, useState, type FormEvent } from "react";
import { Camera, User, Hash, FileText, Save } from "lucide-react";
import type { ProfileResponse } from "../../../app/feature/profile/type";
import { ResolveUrlAsset } from "../../../utils/resolve-url";
import { useImageUploader } from "../../../components/Editor/ImageUploader/ImageProvider";
import { useAlert } from "../../../hooks/useAllert";
import { createProfileApi } from "../../../app/feature/profile";
import { api } from "../../../lib/api/api-client";
import { ApiError } from "../../../lib/api/error";


interface MiniProfileProps {
  profile?: Partial<ProfileResponse>;
}

export  function MiniProfile({
  profile = {},
}: MiniProfileProps) {
    const info = useAlert()
    const imgHandler = useImageUploader()
  const formRef = useRef<HTMLFormElement>(null);
  const [img,setImg] = useState<string>(profile.avatar || "")

  const handleAvatarSubmit = async () => {
    const imgres = await imgHandler.open()
    if (!imgres) {
       await info({title:"Tidak Ada Gambar Yang Dipilih"})
        return
    }
    try{
        const _ = createProfileApi(api).updateAvatar(imgres.url)
        setImg(imgres.url)

    }catch(e){
        if(e instanceof ApiError){
            return await info({title:"Terjadi Kesalahan Code: "+e.code})
        }
        console.log(e)
    }
    
  }

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const fd = new FormData(e.currentTarget);
    const req ={
      username: String(fd.get("username") ?? "").trim(),
      tag: String(fd.get("tag") ?? "").trim().replace(/^#/, ""),
      bio: String(fd.get("bio") ?? "").trim(),
    }
    try{

        const res = await createProfileApi(api).updateProfile(req)
        return info({title:"Profile Berhasil Di Ubah"})
    }catch(e){
        if(e instanceof ApiError) {
            return await info({title:"Terjadi Kesalahan, Code: "+e.code})
        }
        console.log(e)
    }
  };

  const inputWrap =
    "flex items-center gap-3 rounded-full border border-gray-200 bg-white px-4 py-2.5 transition focus-within:border-[#0B63CE] focus-within:ring-2 focus-within:ring-[#0B63CE]/20";
  const inputCls =
    "w-full bg-transparent text-sm text-[#1F2937] placeholder:text-[#6B7280] outline-none";
  const labelCls =
    "mb-1.5 block text-[11px] font-semibold uppercase tracking-widest text-[#6B7280]";

  return (
    <div className="w-full  rounded-2xl border border-gray-100 bg-white p-6 shadow-[0_4px_16px_rgba(31,41,55,0.06)] sm:p-8">
      {/* Eyebrow */}
      <p className="mb-1 text-[11px] font-semibold uppercase tracking-widest text-[#0B63CE]">
        Profil Saya
      </p>
      <h2 className="mb-6 text-xl font-semibold tracking-tight text-[#1F2937]">
        Kelola Profil
      </h2>

      {/* Avatar */}
      <div className="mb-6 flex justify-center">
        <button
          type="button"
          onClick={() => handleAvatarSubmit()}
          aria-label="Ubah foto profil"
          className="group relative h-28 w-28 cursor-pointer overflow-hidden rounded-full bg-[#F5EFE3] ring-4 ring-[#1AA6B7]/20 transition hover:ring-[#1AA6B7]/40 focus-visible:outline-none focus-visible:ring-[#0B63CE]"
        >
          {img != "" ? (
            <img
              src={ResolveUrlAsset(img)}
              alt={`Foto profil ${profile.username ?? "pengguna"}`}
              className="h-full w-full object-cover"
            />
          ) : (
            <span className="flex h-full w-full items-center justify-center text-[#6B7280]">
              <User className="h-10 w-10" />
            </span>
          )}
          <span className="absolute inset-0 flex items-center justify-center bg-[#1F2937]/50 text-white opacity-0 transition group-hover:opacity-100 group-focus-visible:opacity-100">
            <Camera className="h-6 w-6" />
          </span>
        </button>
      </div>

      {/* Form (uncontrolled) */}
      <form ref={formRef} onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="username" className={labelCls}>
            Username
          </label>
          <div className={inputWrap}>
            <User className="h-4 w-4 shrink-0 text-[#0B63CE]" />
            <input
              id="username"
              name="username"
              type="text"
              defaultValue={profile.username}
              placeholder="Nama pengguna"
              required
              className={inputCls}
            />
          </div>
        </div>

        <div>
          <label htmlFor="tag" className={labelCls}>
            Tag
          </label>
          <div className={inputWrap}>
            <Hash className="h-4 w-4 shrink-0 text-[#0B63CE]" />
            <input
              id="tag"
              name="tag"
              type="text"
              defaultValue={profile.tag}
              placeholder="0001"
              required
              className={inputCls}
            />
          </div>
        </div>

        <div>
          <label htmlFor="bio" className={labelCls}>
            Bio
          </label>
          <div className="flex gap-3 rounded-2xl border border-gray-200 bg-white px-4 py-3 transition focus-within:border-[#0B63CE] focus-within:ring-2 focus-within:ring-[#0B63CE]/20">
            <FileText className="mt-0.5 h-4 w-4 shrink-0 text-[#0B63CE]" />
            <textarea
              id="bio"
              name="bio"
              rows={3}
              defaultValue={profile.bio}
              placeholder="Ceritakan sedikit tentang dirimu..."
              className={`${inputCls} resize-none leading-relaxed`}
            />
          </div>
        </div>

        <button
          type="submit"
          className="inline-flex w-full items-center justify-center gap-2 rounded-full bg-[#0B63CE] px-6 py-3 text-sm font-semibold text-white transition hover:bg-[#0952a8] focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#0B63CE] focus-visible:ring-offset-2"
        >
          <Save className="h-4 w-4" />
          Simpan Perubahan
        </button>
      </form>
    </div>
  );
}