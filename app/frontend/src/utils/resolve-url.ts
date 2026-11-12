export function ResolveUrl(imageSource: string, domainResolve: string): string {
    if (/^(https?:|data:)/.test(imageSource)) return imageSource
    return new URL(imageSource, domainResolve).href
}

// src/utils/resolveUrl.ts
const ASSET_PREFIX = "assets:";
const ASSET_BASE_URL =  "/api";

export function ResolveUrlAsset(input?: string | null): string {
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
    return `${base}/assets/${path}`;
  }

  // Sudah public: http(s), //cdn, data:, blob:, atau path relatif (/img/a.png)
  return url;
}