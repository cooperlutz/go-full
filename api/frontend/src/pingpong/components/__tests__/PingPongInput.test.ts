import PingPongInput from "../PingPongInput.vue";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { mount } from "@vue/test-utils";

const mockSendPingPong = vi.fn();
const mockShowResponse = vi.fn();

vi.mock("~/pingpong/composables/usePingPong", () => {
  return {
    useSendPingPong: () => ({ sendPingPong: mockSendPingPong }),
  };
});
vi.mock("~/pingpong/composables/usePingPongToast", () => {
  return {
    useShowCreatePingPongResponse: (msg?: string) => mockShowResponse(msg),
  };
});

beforeEach(() => {
  vi.clearAllMocks();
});

describe("PingPongInput", () => {
  it("renders input and send button", () => {
    const wrapper = mount(PingPongInput);
    expect(wrapper.find("#pingpong-input").exists()).toBe(true);
    expect(wrapper.find("#send-button").exists()).toBe(true);
  });

  it("calls sendPingPong with the input value and shows the response", async () => {
    mockSendPingPong.mockResolvedValueOnce({ message: "pong" });

    const wrapper = mount(PingPongInput);
    const input = wrapper.find("#pingpong-input");
    await input.setValue("ping");
    await wrapper.find("#send-button").trigger("click");
    await Promise.resolve();

    expect(mockSendPingPong).toHaveBeenCalledWith("ping");
    expect(mockShowResponse).toHaveBeenCalledWith("pong");
  });
});
