<template>
  <!-- this part not done by copilot -->
  <div>
    if you're questioning what the deal with this page is and why it is the way
    it is... gh copilot was prompted to build out the entire coming soon page,
    so here it is...
  </div>
  <!-- end -->

  <div class="coming-soon-container">
    <div class="pingpong-logo">
      <svg width="80" height="80" viewBox="0 0 80 80">
        <circle
          cx="40"
          cy="40"
          r="38"
          fill="#00C896"
          stroke="#222"
          stroke-width="4"
        />
        <rect x="30" y="60" width="20" height="10" rx="5" fill="#222" />
        <circle
          cx="55"
          cy="35"
          r="10"
          fill="#fff"
          stroke="#222"
          stroke-width="2"
        />
      </svg>
    </div>
    <h1>Coming Soon</h1>
    <p class="subtitle">
      This page is still warming up its paddle.<br />
      <span class="pingpong-joke"
        >Our devs are busy perfecting their backhand.</span
      >
    </p>
    <div class="pingpong-animation">
      <div class="ball"></div>
      <div class="paddle"></div>
    </div>
    <footer>
      <span>PingPong Web &copy; {{ new Date().getFullYear() }}</span>
    </footer>
  </div>
</template>

<style scoped>
.coming-soon-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #00c896 0%, #00796b 100%);
  color: #222;
  font-family: "Montserrat", "Arial", sans-serif;
  text-align: center;
  position: relative;
}

.pingpong-logo {
  margin-bottom: 1.5rem;
  animation: bounce 1.2s infinite alternate;
}

@keyframes bounce {
  to {
    transform: translateY(-10px);
  }
}

h1 {
  font-size: 3rem;
  font-weight: 900;
  letter-spacing: 2px;
  margin-bottom: 0.5rem;
}

.subtitle {
  font-size: 1.25rem;
  margin-bottom: 2rem;
}

.pingpong-joke {
  display: block;
  margin-top: 0.5rem;
  font-style: italic;
  color: #fff;
  background: rgba(34, 34, 34, 0.1);
  border-radius: 8px;
  padding: 0.25em 0.5em;
}

.pingpong-animation {
  position: relative;
  width: 120px;
  height: 60px;
  margin: 2rem auto;
}

.ball {
  position: absolute;
  left: 10px;
  top: 20px;
  width: 24px;
  height: 24px;
  background: #fff;
  border-radius: 50%;
  border: 2px solid #222;
  animation: ball-move 1.2s infinite cubic-bezier(0.68, -0.55, 0.27, 1.55);
}

@keyframes ball-move {
  0% {
    left: 10px;
    top: 20px;
  }
  50% {
    left: 80px;
    top: 10px;
  }
  100% {
    left: 10px;
    top: 20px;
  }
}

.paddle {
  position: absolute;
  right: 0;
  bottom: 0;
  width: 40px;
  height: 12px;
  background: #222;
  border-radius: 6px;
  transform: rotate(-20deg);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

footer {
  position: absolute;
  bottom: 24px;
  width: 100%;
  text-align: center;
  font-size: 0.95rem;
  color: #fff;
  opacity: 0.8;
}
</style>
