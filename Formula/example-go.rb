class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/archive/refs/tags/example-go-24dcbf0392c3c6d0f48198828a82e85c254eb121.tar.gz"
  version "24dcbf0392c3c6d0f48198828a82e85c254eb121"

  def install
    bin.install "example-go"
    man1.install Dir["*.1"]
    generate_completions_from_executable("./example-go", "completion")
  end
end
