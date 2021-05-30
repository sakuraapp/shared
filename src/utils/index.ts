export function padLeft(str: string, len = 4): string {
    return Array(len - str.length + 1).join('0') + str
}
