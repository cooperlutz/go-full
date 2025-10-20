import { describe, it, vi, beforeEach, expect } from "vitest";
import { useShowCreatePingPongResponse } from "../usePingPongToast";
import Swal from "sweetalert2";

vi.mock("sweetalert2", () => ({
  default: {
    fire: vi.fn(),
  },
}));

describe("useShowCreatePingPongResponse", () => {
  const fireMock = Swal.fire as unknown as ReturnType<typeof vi.fn>;

  beforeEach(() => {
    fireMock.mockClear();
  });

  it("shows notification with default color for undefined", () => {
    useShowCreatePingPongResponse(undefined);
    expect(fireMock).toHaveBeenCalledExactlyOnceWith(
      expect.objectContaining({
        html: "undefined",
        background: "#1084d0",
      }),
    );
  });

  it("shows notification with red background for 'Ping!'", () => {
    useShowCreatePingPongResponse("Ping!");
    expect(fireMock).toHaveBeenCalledExactlyOnceWith(
      expect.objectContaining({
        html: "Ping!",
        background: "#e11d48",
      }),
    );
  });

  it("shows notification with blue background for 'Pong!'", () => {
    useShowCreatePingPongResponse("Pong!");
    expect(fireMock).toHaveBeenCalledExactlyOnceWith(
      expect.objectContaining({
        html: "Pong!",
        background: "#0ea5e9",
      }),
    );
  });

  it("shows notification with default color for other values", () => {
    useShowCreatePingPongResponse("Other");
    expect(fireMock).toHaveBeenCalledExactlyOnceWith(
      expect.objectContaining({
        html: "Other",
        background: "#1084d0",
      }),
    );
  });

  it("sets correct SweetAlert2 options", () => {
    useShowCreatePingPongResponse("Ping!");
    expect(fireMock).toHaveBeenCalledExactlyOnceWith(
      expect.objectContaining({
        position: "bottom-end",
        timer: 1000,
        timerProgressBar: true,
        backdrop: false,
        showConfirmButton: false,
        toast: true,
        animation: false,
        color: "#fff",
      }),
    );
  });
});
