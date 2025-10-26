import PongButton from "../PingPongButtons/PongButton.vue";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { mount } from "@vue/test-utils";

const mockSendPingPong = vi.fn();
const mockShowCreatePingPongResponse = vi.fn();

// Mock composables before importing the component so the component receives the mocks
vi.mock("~/pingpong/composables/usePingPong", () => {
  return {
    useSendPingPong: () => ({ sendPingPong: mockSendPingPong }),
  };
});

vi.mock("~/pingpong/composables/usePingPongToast", () => {
  return {
    useShowCreatePingPongResponse: (msg?: string) =>
      mockShowCreatePingPongResponse(msg),
  };
});

beforeEach(() => {
  vi.clearAllMocks();
});

describe("PongButton", () => {
  it('renders a button with text "pong"', () => {
    const wrapper = mount(PongButton);
    expect(wrapper.text()).toContain("pong");
  });

  it("calls sendPingPong with 'pong' and shows the response message on click", async () => {
    mockSendPingPong.mockResolvedValue({ message: "pong" });

    const wrapper = mount(PongButton);
    await wrapper.find("#pong-button").trigger("click");

    // wait for the async handler to resolve
    await Promise.resolve();

    expect(mockSendPingPong).toHaveBeenCalledWith("pong");
    expect(mockShowCreatePingPongResponse).toHaveBeenCalledWith("pong");
  });

  it("calls show response with undefined when sendPingPong returns undefined", async () => {
    mockSendPingPong.mockResolvedValue(undefined);

    const wrapper = mount(PongButton);
    await wrapper.find("#pong-button").trigger("click");

    await Promise.resolve();

    expect(mockSendPingPong).toHaveBeenCalledWith("pong");
    expect(mockShowCreatePingPongResponse).toHaveBeenCalledWith(undefined);
  });
});
