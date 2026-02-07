import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";
import PingPongTable from "../PingPongTable.vue";

describe("PingPongTable", () => {
  it("renders card with title and description", () => {
    // Arrange & Act
    const wrapper = mount(PingPongTable);

    // Assert
    expect(wrapper.find("h2").text()).toBe("Ping Pongs");
    expect(wrapper.find("p").text()).toBe("All Ping Pongs");
  });

  it("shows table when data is loaded", async () => {
    // Arrange
    const wrapper = mount(PingPongTable);
    await nextTick();

    // Act
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    expect(wrapper.find("table").exists()).toBe(true);
  });

  it("renders ping pong data in table rows", async () => {
    // Arrange
    const wrapper = mount(PingPongTable);
    await nextTick();

    // Act
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    const headers = wrapper.findAll("th");
    expect(headers).toHaveLength(6);
    expect(headers[0]?.text()).toBe("Ping Pong ID");
    expect(headers[1]?.text()).toBe("Message");
    expect(headers[2]?.text()).toBe("Created At");
    expect(headers[3]?.text()).toBe("Updated At");
    expect(headers[4]?.text()).toBe("Deleted At");
    expect(headers[5]?.text()).toBe("Deleted");

    const rows = wrapper.findAll("tbody tr");
    expect(rows.length).toBeGreaterThan(0);
    const firstRowCells = rows[0]?.findAll("td");
    expect(firstRowCells?.length).toBe(6);
    expect(firstRowCells?.[0]?.text()).toBe(
      "f660452b-4075-4eac-b87a-a5b1ce7bd428",
    );
    expect(firstRowCells?.[1]?.text()).toBe("pong");
    expect(firstRowCells?.[4]?.text()).toBe("");
    expect(firstRowCells?.[5]?.text()).toBe("false");
  });

  it("shows loading state initially", async () => {
    // Arrange & Act
    const wrapper = mount(PingPongTable);
    await nextTick();

    // Assert
    expect(wrapper.text()).toContain("Loading ping pongs...");
  });

  it("shows error message on fetch failure", async () => {
    // Arrange
    global.fetch = vi.fn(() =>
      Promise.resolve({
        ok: false,
        statusText: "Internal Server Error",
      } as Response),
    ) as unknown as typeof fetch;

    const wrapper = mount(PingPongTable);

    // Act
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    expect(wrapper.text()).toContain("Error loading ping pongs:");
  });
});
