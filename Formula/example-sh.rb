class ExampleSh < Formula
  desc "A simple shell example."
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-sh-e55fb463745657a6a305ef7980cbf16fa1e76750/example-sh"
  version "e55fb463745657a6a305ef7980cbf16fa1e76750"

  def install
    bin.install "example-sh"
  end
end
