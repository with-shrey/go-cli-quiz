package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

func TestBooks(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Quiz-CLI")
}

var testData1 [][]string = [][]string{
	{
		"1", "2",
	},
	{
		"3", "4",
	},
	{
		"5", "6",
	},
	{
		"7", "8",
	},
}


var testData2 [][]string = [][]string{
	{
		"9", "10",
	},
	{
		"11", "12",
	},
	{
		"13", "14",
	},
	{
		"15", "16",
	},

}

func buildFakeCsv() string {
	dir, err := ioutil.TempDir("../../", "temp")
	if err != nil {
		panic(err)
	}
	csvFile1, _ := os.Create(filepath.Join(dir,  "test1.csv"))
	defer csvFile1.Close()
	csvFile2, _ := os.Create(filepath.Join(dir,  "test2.csv"))
	defer csvFile2.Close()
	csvFile3, _ := os.Create("problems.csv")
	defer csvFile3.Close()
	csvwriter1 := csv.NewWriter(csvFile1)
	csvwriter1.WriteAll(testData1);
	csvwriter2 := csv.NewWriter(csvFile2)
	csvwriter2.WriteAll(testData1);
	csvwriter3 := csv.NewWriter(csvFile3)
	csvwriter3.WriteAll(testData2)
	return dir
}

func buildCLI() string {
	cliPath, err := gexec.Build("github.com/with-shrey/go-quiz/cmd/cli", )
	Expect(err).NotTo(HaveOccurred())
	return cliPath
}

func runCLI(path string, args ...string) (*gexec.Session, io.WriteCloser) {
	cmd := exec.Command(path, args...)
	stdin, _ := cmd.StdinPipe()
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session, stdin
}

func getQuestionsText(testData [][]string) string {
	questions := []string {}
	for _, row := range testData {
		questions = append(questions, "Question: " + row[0])
	}
	return strings.Join(questions, "\n")
}

var _ = Describe("Quiz-CLI", func() {
	var session *gexec.Session
	var scratchpadPath string
	var tempDir string
	var stdin io.WriteCloser

	BeforeEach(func() {
		tempDir = buildFakeCsv()
		scratchpadPath = buildCLI()
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
		os.Remove("problems.csv")
		gexec.CleanupBuildArtifacts()
	})

  Describe("Question imports", func() {
    Context("without any flag ", func() {
      It("questions should be from problems.csv", func() {
				session, stdin = runCLI(scratchpadPath)
				stdin.Write([]byte("\n\n\n\n\n\n"))
				Eventually(session).Should(gbytes.Say(getQuestionsText(testData2)))
      })
			
			It("timeout should be default", func() {
				session, stdin = runCLI(scratchpadPath)
				timeout, _ := strconv.Atoi(DefaultTimeout)
				session.Wait(timeout + 1)
				Eventually(session).Should(gbytes.Say("Time Out"))
      })
    })

		Context("with -csv flag ", func() {
      It("first question should be from test1.csv", func() {
				session, stdin = runCLI(scratchpadPath, "-csv", filepath.Join(tempDir, "test1.csv"))
				stdin.Write([]byte("\n\n\n\n\n\n"))
				Eventually(session).Should(gbytes.Say(getQuestionsText(testData1)))
      })
			
			It("timeout should be default", func() {
				session, stdin = runCLI(scratchpadPath, "-timeout", DefaultTimeout)
				timeout, _ := strconv.Atoi(DefaultTimeout)
				session.Wait(timeout + 1)
				Eventually(session).Should(gbytes.Say("Time Out"))
      })
    })
		
		Context("with -timeout flag only", func() {
      It("should print questions from problems.csv", func() {
				session, stdin = runCLI(scratchpadPath, "-timeout", "5")
				stdin.Write([]byte("\n\n\n\n\n\n"))
				Eventually(session).Should(gbytes.Say(getQuestionsText(testData2)))
      })
			
			It("should timeout after specified value", func() {
				session, stdin = runCLI(scratchpadPath, "-timeout", "5")
				session.Wait(5 + 1)
				Eventually(session).Should(gbytes.Say("Time Out"))
      })
    })
		
		Context("after timeout happens", func() {
			It("should print correct answer and total questions", func() {
				session, stdin = runCLI(scratchpadPath, "-timeout", "5")
				stdin.Write([]byte("10\n12\n0\n0\n"))
				session.Wait(5 + 1)
				Eventually(session).Should(gbytes.Say("Correct Answers: 2\nTotal Question: 4"))
			})
		})
		
		Context("after quiz is complete", func() {
			It("should print correct answer and total questions", func() {
				session, stdin = runCLI(scratchpadPath)
				stdin.Write([]byte("10\n12\n0\n0\n"))
				Eventually(session).Should(gbytes.Say("Correct Answers: 2\nTotal Question: 4"))
			})
		})
  })
})
