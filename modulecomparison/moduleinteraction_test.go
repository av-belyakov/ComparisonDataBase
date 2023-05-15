package modulecomparison_test

import (
	"fmt"
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/av-belyakov/comparisondatabase/commonlibs"
)

var _ = Describe("Moduleinteraction", func() {
	Context("Тест 1. Просто тест математических вычислений", func() {
		It("Должно быть выполненно деление с округлением в большую сторону", func() {
			num := math.Round(542 / 150)

			fmt.Println("RESULT = ", num)

			Expect(num).Should(Equal(4))
		})

		It("Должно быть полученно заданное количество частей", func() {
			Expect(commonlibs.GetCountChunk(int64(541), 150)).Should(Equal(4))
		})
	})
})
