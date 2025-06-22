class ExamplePy < Formula
  desc "A simple Python example."
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-py-e55fb463745657a6a305ef7980cbf16fa1e76750/example-py"
  version "e55fb463745657a6a305ef7980cbf16fa1e76750"

  depends_on "python@3.12"

  def install
    bin.install "example-py"
  end
end
