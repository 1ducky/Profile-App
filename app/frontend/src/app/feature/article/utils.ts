// src/utils/resolveUrl.ts
const ASSET_PREFIX = "asset:";
const ASSET_BASE_URL = import.meta.env.PUBLIC_API_URL ?? "";

export function ResolveUrl(input?: string | null): string {
  if (!input) return "";

  const url = input.trim();

  // Format internal: asset:post/filekey.ext
  if (url.startsWith(ASSET_PREFIX)) {
    const path = url
      .slice(ASSET_PREFIX.length)
      .replace(/^\/+/, "") // buang slash di depan
      .split("/")
      .map(encodeURIComponent) // encode tiap segmen, slash tetap aman
      .join("/");

    const base = ASSET_BASE_URL.replace(/\/+$/, "");
    return `${base}/${path}`;
  }

  // Sudah public: http(s), //cdn, data:, blob:, atau path relatif (/img/a.png)
  return url;
}