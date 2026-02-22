import { describe, it, expect, vi } from "vitest";
import { mount } from "@vue/test-utils";
import { nextTick } from "vue";
import ExamLibraryTable from "../ExamLibraryTable.vue";

describe("ExamLibraryTable", () => {
  it("shows table when data is loaded", async () => {
    // Arrange
    const wrapper = mount(ExamLibraryTable);
    await nextTick();

    // Act
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    expect(wrapper.find("table").exists()).toBe(true);
  });

  it("renders exam data in table rows", async () => {
    // Arrange
    const wrapper = mount(ExamLibraryTable);
    await nextTick();

    // Act
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    const headers = wrapper.findAll("th");
    expect(headers).toHaveLength(4);
    expect(headers[0]?.text()).toBe("Exam ID");
    expect(headers[1]?.text()).toBe("Name");
    expect(headers[2]?.text()).toBe("Grade Level");
    expect(headers[3]?.text()).toBe("Time Limit (seconds)");

    const rows = wrapper.findAll("tbody tr");
    expect(rows.length).toBeGreaterThan(0);
    const firstRowCells = rows[0]?.findAll("td");
    expect(firstRowCells?.length).toBe(4);
    expect(firstRowCells?.[0]?.text()).toBe(
      "5d9abb80-0706-42ad-8131-33627d3e6b17",
    );
    expect(firstRowCells?.[1]?.text()).toBe("Midterm Exam");
    expect(firstRowCells?.[2]?.text()).toBe("3");
    expect(firstRowCells?.[3]?.text()).toBe("3600");
  });

  it("shows loading state initially", async () => {
    // Arrange & Act
    const wrapper = mount(ExamLibraryTable);
    await nextTick();

    // Assert
    expect(wrapper.find("#exam-table-loading").exists()).toBe(true);
  });

  it("shows error message on fetch failure", async () => {
    // Arrange
    global.fetch = vi.fn(() => Promise.reject(new Error("ruh roh")));

    const wrapper = mount(ExamLibraryTable);

    // Act
    await new Promise((resolve) => setTimeout(resolve, 100));
    await nextTick();

    // Assert
    expect(wrapper.text()).toContain(
      "FetchError: The request failed and the interceptors did not return an alternative response",
    );
  });
});
