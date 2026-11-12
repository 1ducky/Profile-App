import {
    useRef,
    useState,
    type KeyboardEvent,
    type MouseEvent,
    type PointerEvent,
} from "react";

/**
 * Logika slider geser (pointer drag) yang bisa dipakai ulang.
 * Pakai: viewport (overflow-hidden) > track (flex, style transform) > slide (basis-full).
 */
export function useDragSlider(count: number) {
    const viewportRef = useRef<HTMLDivElement>(null);
    const [index, setIndex] = useState(0);
    const [dx, setDx] = useState(0);
    const [dragging, setDragging] = useState(false);
    const drag = useRef({ startX: 0, startT: 0, dx: 0, active: false, moved: false });

    const go = (i: number) => setIndex(Math.min(Math.max(i, 0), count - 1));

    const onPointerDown = (e: PointerEvent<HTMLDivElement>) => {
        if (e.pointerType === "mouse" && e.button !== 0) return;
        drag.current = {
            startX: e.clientX,
            startT: performance.now(),
            dx: 0,
            active: true,
            moved: false,
        };
    };

    const onPointerMove = (e: PointerEvent<HTMLDivElement>) => {
        const d = drag.current;
        if (!d.active) return;
        let delta = e.clientX - d.startX;

        // Ambang kecil agar klik biasa tidak dianggap geseran
        if (!d.moved) {
            if (Math.abs(delta) < 6) return;
            d.moved = true;
            setDragging(true);
            e.currentTarget.setPointerCapture(e.pointerId);
        }

        // Tertahan di ujung
        if ((index === 0 && delta > 0) || (index === count - 1 && delta < 0)) {
            delta *= 0.35;
        }
        d.dx = delta;
        setDx(delta);
    };

    const endDrag = () => {
        const d = drag.current;
        if (!d.active) return;
        d.active = false;

        if (d.moved) {
            const width = viewportRef.current?.offsetWidth ?? 1;
            const velocity = d.dx / Math.max(performance.now() - d.startT, 1); // px/ms
            if (d.dx < -width * 0.2 || velocity < -0.5) go(index + 1);
            else if (d.dx > width * 0.2 || velocity > 0.5) go(index - 1);
        }
        setDragging(false);
        setDx(0);
    };

    const onClickCapture = (e: MouseEvent) => {
        if (drag.current.moved) {
            e.preventDefault();
            e.stopPropagation();
            drag.current.moved = false;
        }
    };

    const onKeyDown = (e: KeyboardEvent<HTMLElement>) => {
        if (e.key === "ArrowLeft") go(index - 1);
        if (e.key === "ArrowRight") go(index + 1);
    };

    return {
        viewportRef,
        index,
        go,
        dragging,
        onKeyDown,
        trackProps: {
            onPointerDown,
            onPointerMove,
            onPointerUp: endDrag,
            onPointerCancel: endDrag,
            onClickCapture,
            style: { transform: `translate3d(calc(${-index * 100}% + ${dx}px), 0, 0)` },
        },
    };
}