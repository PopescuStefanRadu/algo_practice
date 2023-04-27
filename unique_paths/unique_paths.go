package unique_paths

import "fmt"

func uniquePaths(m int, n int) int {
	return UniquePaths(m, n, map[string]int{})
}

func UniquePaths(m, n int, memory map[string]int) int {
	if m == 1 || n == 1 {
		return 1
	}
	identifier := fmt.Sprintf("%v-%v", m, n)
	if memorized, exists := memory[identifier]; exists {
		return memorized
	}

	memory[identifier] = UniquePaths(m-1, n, memory) + UniquePaths(m, n-1, memory)
	return memory[identifier]
}
