import {useEffect, useRef} from "react";

export function useInterval(callback: Function, delay: number) {
    const savedCallback = useRef<Function>();

    useEffect(() => {
        savedCallback.current = callback
    }, [callback]);

    useEffect(() => {
        function tick() {
            if (savedCallback.current) {
                savedCallback?.current();
            }
        }

        const id = setInterval(tick, delay);
        return () => {
            clearInterval(id);
        }
    })
}
