import pingpongHandlers from "./ping-pongs";
import metricsHandlers from "./metrics";

export const pingPongMockHandlers = [...pingpongHandlers, ...metricsHandlers];
