import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import App from "../App.vue";

describe("App.vue", () => {
  it("renders router-view and passes the component", () => {
    const RouterViewStub = {
      template: '<div class="router-view-stub"></div>',
    };
    const wrapper = mount(App, {
      global: {
        stubs: {
          RouterView: RouterViewStub,
        },
      },
    });
    expect(wrapper.find(".router-view-stub").exists()).toBe(true);
  });
});
