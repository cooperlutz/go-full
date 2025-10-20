async function simulateDelay(): Promise<void> {
  await new Promise((resolve) => setTimeout(resolve, 1000)); // Simulate delay
}

export { simulateDelay };
