import PingButton from "../PingPongButtons/PingButton.vue";
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

describe("PingButton", () => {
  it('renders a button with text "ping"', () => {
    const wrapper = mount(PingButton);
    expect(wrapper.text()).toContain("ping");
  });

  it("calls sendPingPong with 'ping' and shows the response message on click", async () => {
    mockSendPingPong.mockResolvedValue({ message: "pong" });

    const wrapper = mount(PingButton);
    await wrapper.find("#ping-button").trigger("click");

    // wait for the async handler to resolve
    await Promise.resolve();

    expect(mockSendPingPong).toHaveBeenCalledWith("ping");
    expect(mockShowCreatePingPongResponse).toHaveBeenCalledWith("pong");
  });

  it("calls show response with undefined when sendPingPong returns undefined", async () => {
    mockSendPingPong.mockResolvedValue(undefined);

    const wrapper = mount(PingButton);
    await wrapper.find("#ping-button").trigger("click");

    await Promise.resolve();

    expect(mockSendPingPong).toHaveBeenCalledWith("ping");
    expect(mockShowCreatePingPongResponse).toHaveBeenCalledWith(undefined);
  });
});
