/**
 * Helper & type generik untuk chart. Tidak ada istilah domain di sini:
 * chart hanya tahu "label sumbu-x" dan "satu atau lebih seri angka".
 */

export interface ChartSeries {
    /** Pengenal unik seri */
    key: string;
    /** Teks di legenda dan tooltip */
    label: string;
    /** Satu nilai per label. `null` = tidak ada data (garis terputus / bar kosong) */
    values: (number | null)[];
    /** Hex. Jika kosong dipakai PALETTE sesuai urutan */
    color?: string;
}

export interface ResolvedSeries extends ChartSeries { color: string }

/** Biru (primary), teal, hijau, pasir — dari design.md */
export const PALETTE = ['#0B63CE', '#1AA6B7', '#16A34A', '#B98A3E'];

export const resolveSeries = (series: ChartSeries[]): ResolvedSeries[] =>
    series.map((s, i) => ({ ...s, color: s.color ?? PALETTE[i % PALETTE.length] }));

export const hasData = (labels: string[], series: ChartSeries[]) =>
    labels.length > 0 && series.some((s) => s.values.some((v) => v !== null && v !== undefined));

const idFull = new Intl.NumberFormat('id-ID');
const idCompact = new Intl.NumberFormat('id-ID', { notation: 'compact', maximumFractionDigits: 1 });
export const formatFull = (n: number) => idFull.format(n);
export const formatCompact = (n: number) => idCompact.format(n);

/** Skala "bulat": 0 … max dengan ±count interval */
export function niceScale(maxValue: number, count = 4) {
    if (!isFinite(maxValue) || maxValue <= 0) return { max: 1, ticks: [0, 1] };
    const raw = maxValue / count;
    const pow = 10 ** Math.floor(Math.log10(raw));
    const f = raw / pow;
    const step = (f <= 1 ? 1 : f <= 2 ? 2 : f <= 2.5 ? 2.5 : f <= 5 ? 5 : 10) * pow;
    const intervals = Math.ceil(maxValue / step);
    const ticks = Array.from({ length: intervals + 1 }, (_, i) => +(i * step).toFixed(10));
    return { max: ticks[ticks.length - 1], ticks };
}

/** Label sumbu-x mana yang ditampilkan supaya tidak saling menimpa */
export function visibleLabelIndexes(n: number, maxLabels: number) {
    const step = Math.max(1, Math.ceil(n / Math.max(1, maxLabels)));
    return Array.from({ length: n }, (_, i) => i).filter((i) => i % step === 0);
}

/** Posisi x (0–100) tiap label */
export const xPositions = (n: number, kind: 'line' | 'bar') =>
    Array.from({ length: n }, (_, i) =>
        kind === 'bar' ? ((i + 0.5) / n) * 100 : n === 1 ? 50 : (i / (n - 1)) * 100,
    );