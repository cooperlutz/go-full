import { useLocalTokenStore } from "~/iam/stores/useToken";

const tokenStore = useLocalTokenStore();

export const getAuthorizationHeader = `Bearer ${tokenStore.getAccessToken()}`;
