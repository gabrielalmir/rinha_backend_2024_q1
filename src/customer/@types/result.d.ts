export type Result<T> = {
    success: boolean;
    data: T;
    error?: any;
};
