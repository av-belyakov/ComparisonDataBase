package logging_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/av-belyakov/comparisondatabase/logging"
)

var _ = Describe("Logging", Ordered, func() {
	listTypeFile := []string{"error", "information", "test_1", "test_2"}
	ld, err := logging.NewLoggingData("logs_app", "/home/artemij/go/src/comparison_data_two_DB", listTypeFile)

	Context("Тест 1. Проверка функции 'NewLoggingData'", func() {
		It("При проверки функции 'NewLoggingData' должны быть созданы директория с лог файлами и сами лог файлы заданных типов", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ld.GetCountFileDescription()).Should(Equal(len(listTypeFile)))
		})
	})

	Context("Тест 2. Получаем список типов лог файлов", func() {
		It("При проверки функции 'GetListTypeFile' должен быть получен список типов лог-файлов и одно из указанных имен", func() {
			list := ld.GetListTypeFile()

			var isExist bool
			for i := 0; i < len(list); i++ {
				if list[i] == "error" {
					isExist = true

					break
				}
			}

			Expect(len(list)).Should(Equal(len(listTypeFile)))
			Expect(isExist).Should(BeTrue())
		})
	})

	Context("Тест 3. Записываем информацию в лог файлы error.log и information.log", func() {
		It("При проверки функции 'WriteLoggingData' должны быть успешно записаны несколько строк в файлы error.log и information.log", func() {
			testList := map[string]string{
				"error":       "new test error message",
				"information": "new my test information message"}

			for k, v := range testList {
				isSuccess := ld.WriteLoggingData(v, k)

				Expect(isSuccess).Should(BeTrue())
			}
		})
	})

	AfterAll(func() {
		ld.ClosingFiles()
	})
})
