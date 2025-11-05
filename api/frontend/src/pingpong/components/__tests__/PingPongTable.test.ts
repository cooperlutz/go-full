import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";
import PingPongTable from "../PingPongTable.vue";

describe("PingPongTable", () => {
  it("renders card with title and description", () => {
    const wrapper = mount(PingPongTable);

    expect(wrapper.find("h2").text()).toBe("Ping Pongs");
    expect(wrapper.find("p").text()).toBe("All Ping Pongs");
  });

  it("renders table headers correctly", async () => {
    const wrapper = mount(PingPongTable);
    await nextTick();

    const headers = wrapper.findAll("th");
    expect(headers).toHaveLength(6);
    expect(headers[0]?.text()).toBe("Ping Pong ID");
    expect(headers[1]?.text()).toBe("Message");
    expect(headers[2]?.text()).toBe("Created At");
    expect(headers[3]?.text()).toBe("Updated At");
    expect(headers[4]?.text()).toBe("Deleted At");
    expect(headers[5]?.text()).toBe("Deleted");
  });

  it("shows table when data is loaded", async () => {
    const wrapper = mount(PingPongTable);

    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    expect(wrapper.find("table").exists()).toBe(true);
  });

  it("renders ping pong data in table rows", async () => {
    const wrapper = mount(PingPongTable);
    await nextTick();

    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    const rows = wrapper.findAll("tbody tr");
    expect(rows.length).toBeGreaterThan(0);

    if (rows.length > 0) {
      const firstRowCells = rows[0]?.findAll("td");
      expect(firstRowCells?.length).toBe(6);
      expect(firstRowCells?.[0]?.text()).toBe(
        "f660452b-4075-4eac-b87a-a5b1ce7bd428",
      );
      expect(firstRowCells?.[1]?.text()).toBe("pong");
      expect(firstRowCells?.[2]?.text()).toBe(
        "Thu Oct 30 2025 08:19:44 GMT-0500 (Central Daylight Time)",
      );
      expect(firstRowCells?.[3]?.text()).toBe(
        "Sun Dec 31 0000 18:09:24 GMT-0550 (Central Standard Time)",
      );
      expect(firstRowCells?.[4]?.text()).toBe("");
      expect(firstRowCells?.[5]?.text()).toBe("false");
    }
  });
});
