class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-e55fb463745657a6a305ef7980cbf16fa1e76750/example-go"
  version "e55fb463745657a6a305ef7980cbf16fa1e76750"

  def install
    bin.install "example-go"

    system bin/"example-go", "man"
    man1.install Dir["*.1"]

    generate_completions_from_executable(bin/"example-go", "completion")
  end
end
