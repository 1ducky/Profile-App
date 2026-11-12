// ============================================================
// C. MOCK: STATS
// ============================================================

import type { ArticleStats } from "../../types/article/entity";

export const mockStats: Record<string, ArticleStats> = {
    viral: { views: 128_400, comments: 342, shares: 1_210 },
    medium: { views: 12_800, comments: 45, shares: 87 },
    fresh: { views: 1_240, comments: 3, shares: 5 },
};
