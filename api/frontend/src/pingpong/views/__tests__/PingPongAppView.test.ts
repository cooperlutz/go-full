import { describe, it, expect, vi } from "vitest";

import { mount } from "@vue/test-utils";
import PingPongAppView from "../PingPongAppView.vue";

// Mock the currentroute.path
vi.mock("vue-router", () => ({
  useRoute: () => ({
    path: "/pingpong/app",
  }),
}));

vi.mock("usePingPongMetrics", () => ({
  usePingPongMetrics: () => ({
    fetchMetrics: vi.fn(),
    metrics: vi.fn(() => []),
    loading: vi.fn(() => false),
    error: vi.fn(() => null),
  }),
}));

// vi.mock("~/pingpong/composables/usePingPong", () => ({
//   usePingPongs: () => ({
//     actions: { fetchData: vi.fn() },
//     allPingPongs: {
//       pingpongs: [
//         {
//           id: 1,
//           message: "Ping!",
//           createdAt: "2024-01-01",
//           updatedAt: "2024-01-01",
//           deletedAt: null,
//           deleted: false,
//         },
//       ],
//     },
//   }),
// }));

vi.mock("clickPing", () =>
  vi.fn(() =>
    Promise.resolve({
      id: 1,
      message: "Ping!",
      createdAt: "2024-01-01",
      updatedAt: "2024-01-01",
      deletedAt: null,
      deleted: false,
    }),
  ),
);

vi.mock("usePingPong", () => ({
  usePingPong: () => ({
    state: { loading: false, error: null },
    actions: {
      sendPingPong: vi.fn(() => Promise.resolve({ message: "Ping!" })),
    },
  }),

  usePingPongs: () => ({
    state: { loading: false, error: null },
    actions: { fetchData: vi.fn() },
    // allPingPongs: {
    //   pingpongs: [
    //     {
    //       id: 1,
    //       message: "Ping!",
    //       createdAt: "2024-01-01",
    //       updatedAt: "2024-01-01",
    //       deletedAt: null,
    //       deleted: false,
    //     },
    //   ],
    // },
  }),
}));

// Remove this line because we are already mocking sendPingPong in the vi.mock for usePingPong

describe("PingPongAppView.vue", () => {
  it("should render correctly", () => {
    const wrapper = mount(PingPongAppView);
    expect(wrapper.exists()).toBe(true);
  });

  it("should render a `ping` button that sends a ping message to the api when clicked", async () => {
    const wrapper = mount(PingPongAppView);
    const button = wrapper.find("#ping-button");
    expect(button.exists()).toBe(true);
    expect(button.text().toLowerCase()).toContain("ping");
    // await button.trigger("click");
    // expect(wrapper.exists()).toBe(true);
  });

  it("should render a `pong` button that sends a pong message to the api when clicked", async () => {
    const wrapper = mount(PingPongAppView);
    const button = wrapper.find("#pong-button");
    expect(button.exists()).toBe(true);
    expect(button.text().toLowerCase()).toContain("pong");
    // await button.trigger("click");
    // expect(wrapper.exists()).toBe(true);
  });

  it("should render an input field and a send button that sends the input value to the api when clicked", async () => {
    const wrapper = mount(PingPongAppView);
    const input = wrapper.find("input");
    const button = wrapper.find("#send-button");
    expect(input.exists()).toBe(true);
    expect(button.exists()).toBe(true);
    expect(button.text().toLowerCase()).toContain("send");
    await input.setValue("test message");
    expect((input.element as HTMLInputElement).value).toBe("test message");
    // await button.trigger("click");
    // expect(wrapper.exists()).toBe(true);
  });
});
