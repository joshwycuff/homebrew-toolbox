class ExampleSh < Formula
  desc "A simple shell example."
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-sh-b185f0ccd960959ee20ad2c2288a06d01da81191/example-sh"
  version "b185f0ccd960959ee20ad2c2288a06d01da81191"

  def install
    bin.install "example-sh"
  end
end
