class ExamplePy < Formula
  desc "A simple Python example."
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-py-b185f0ccd960959ee20ad2c2288a06d01da81191/example-py"
  version "b185f0ccd960959ee20ad2c2288a06d01da81191"

  depends_on "python@3.12"

  def install
    bin.install "example-py"
  end
end
