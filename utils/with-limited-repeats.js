function withLimitedRepeats(
  promiseFn,
  defaultRes,
  retryLimit,
  refreshRepeatsAfterMs,
) {
  let retryCount = 0;
  let timerId = null;
  return async (...args) => {
    if (retryCount > retryLimit - 1) return defaultRes;
    if (refreshRepeatsAfterMs) {
      if (timerId) clearTimeout(timerId);
      timerId = setTimeout(() => {
        retryCount = 0;
        timerId = null;
      }, refreshRepeatsAfterMs);
    }
    const res = await promiseFn(...args);
    retryCount++;
    return res;
  };
}

const fn = withLimitedRepeats(
  () => Promise.resolve("result"),
  "default",
  4,
  1_500,
);
(async () => {
  for (let i = 0; i < 10; i++) {
    console.log(await fn());
    await new Promise((res) => setTimeout(res, 495));
  }
})();
