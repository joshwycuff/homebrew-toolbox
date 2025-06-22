class ExampleSh < Formula
  desc "A simple shell example."
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-sh-07e30d1afcfd25b65c6da110a4f9c53bd204e3fc/example-sh"
  version "07e30d1afcfd25b65c6da110a4f9c53bd204e3fc"

  def install
    bin.install "example-sh"
  end
end
