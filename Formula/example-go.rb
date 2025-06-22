class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-46cf4ee9e16a74d35027fa4411ba0a8855c14ccc/example-go"
  version "46cf4ee9e16a74d35027fa4411ba0a8855c14ccc"

  def install
    bin.install "example-go"

    system bin/"example-go", "man"
    man1.install Dir["*.1"]

    generate_completions_from_executable(bin/"example-go", "completion")
  end
end
