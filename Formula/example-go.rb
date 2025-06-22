class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-e03116d907ae8e2538d31530f5e47d0a568e1098/example-go"
  version "e03116d907ae8e2538d31530f5e47d0a568e1098"

  def install
    bin.install "example-go"
    system "xattr", "-d", "com.apple.quarantine", "#{bin}/example-go"

    system bin/"example-go", "man"
    man1.install Dir["*.1"]

    generate_completions_from_executable(bin/"example-go", "completion")
  end
end
